#include "notification.h"
#include "map.h"
#include <stdlib.h>
#include <string.h>

Notification *create_notification(char const *id, float score,
                                  float probability, long timestamp) {

  Notification *notification = (Notification *)malloc(sizeof(Notification));

  if (notification == NULL) {
    return NULL;
  }

  notification->id = strdup(id);
  notification->probability = probability;
  notification->score = score;
  notification->timestamp = timestamp;

  // Setup the compute sums for later
  notification->m_plus = 0;
  notification->m_plus_count = 0;
  notification->m_plus_sum = 0;

  notification->m_minus = 1.0; // Avoid division by zero later
  notification->m_minus_count = 0;
  notification->m_minus_sum = 0;
  notification->selected = 0;

  return notification;
}

void free_notification(Notification *notification) {
  if (notification->id != NULL) {
    free(notification->id);
    notification->id = NULL;
  }

  free(notification);
}

NotificationArray *create_notification_list(long length) {

  Notification **notifications =
      (Notification **)calloc(length, sizeof(Notification *));

  if (notifications == NULL) {
    return NULL;
  }

  for (size_t i = 0; i < length; i++) {
    notifications[i] = NULL;
  }

  NotificationArray *array =
      (NotificationArray *)malloc(sizeof(NotificationArray));
  if (array == NULL) {
    free(notifications);

    return NULL;
  }
  array->length = length;
  array->array = notifications;
  array->current_id = 0;

  return array;
}

void free_notification_list(NotificationArray *notifications) {
  if (notifications != NULL) {
    for (size_t i = 0; i < notifications->length; i++) {
      free_notification(notifications->array[i]);
      notifications->array[i] = NULL;
    }

    free(notifications);
  }
}

NotificationArray *parse_python_notification_list(PyObject *list) {
  Py_ssize_t length = PyList_Size(list);

  NotificationArray *array = create_notification_list(length);
  if (array == NULL) {
    PyErr_NoMemory();

    return NULL;
  }

  for (Py_ssize_t i = 0; i < length; i++) {
    PyObject *obj = PyList_GetItem(list, i);

    if (!PyDict_Check(obj)) {
      free_notification_list(array);
      PyErr_Format(PyExc_ValueError, "List item zd% is not an object!", obj);

      return NULL;
    }

    // Parse Items
    PyObject *id_obj = PyDict_GetItemString(obj, "id");
    PyObject *score_obj = PyDict_GetItemString(obj, "score");
    PyObject *probability_obj = PyDict_GetItemString(obj, "probability");
    PyObject *timestamp_obj = PyDict_GetItemString(obj, "timestamp");

    if (!id_obj || !score_obj || !probability_obj || !timestamp_obj) {
      PyErr_SetString(PyExc_ValueError, "Fields not present");

      return NULL;
    }

    const char *id = PyUnicode_AsUTF8(id_obj);
    float score = PyFloat_AsDouble(score_obj);
    float probability = PyFloat_AsDouble(probability_obj);
    long timestamp = PyLong_AsLong(timestamp_obj);

    Notification *notification =
        create_notification(id, score, probability, timestamp);
    if (notification == NULL) {
      free_notification_list(array);

      return NULL;
    }

    array->array[i] = notification;
  }

  return array;
}

hashmap *create_notification_map_from_list(NotificationArray *array) {
  hashmap *notifications = hashmap_create();
  if (notifications == NULL) {
    return NULL;
  }

  // Add the notifications over
  for (size_t i = 0; i < array->length; i++) {
    Notification *notification = array->array[i];

    hashmap_set(notifications, notification->id, strlen(notification->id),
                (uintptr_t)notification);
  }

  return notifications;
}

PyObject *notification_list_to_python_list(NotificationArray *array) {
  PyObject *list = PyList_New(array->length);
  if (list == NULL) {
    PyErr_NoMemory();

    return NULL;
  }

  // Notification Array
  for (size_t i = 0; i < array->length; i++) {
    Notification *notification = array->array[i];

    PyObject *dict = PyDict_New();
    if (dict == NULL) {
      PyErr_NoMemory();
      Py_DECREF(list);

      return NULL;
    }

    PyObject *id_obj = PyUnicode_FromString(notification->id);
    PyObject *score_obj = PyFloat_FromDouble(notification->score);
    PyObject *probability_obj = PyFloat_FromDouble(notification->probability);
    PyObject *selected_obj = PyBool_FromLong(notification->selected);

    if (!id_obj || !score_obj || !probability_obj || !selected_obj) {
      PyErr_NoMemory();
      Py_DECREF(list);
      Py_DECREF(dict);

      return NULL;
    }

    PyDict_SetItemString(dict, "id", id_obj);
    PyDict_SetItemString(dict, "score", score_obj);
    PyDict_SetItemString(dict, "probability", probability_obj);
    PyDict_SetItemString(dict, "selected", selected_obj);

    PyList_SetItem(list, i, dict);
  }

  return list;
}

void print_notification_list(NotificationArray *list) {
  printf("\n--- Notification List ---\n");
  for (size_t i = 0; i < list->length; i++) {
    Notification *notification = list->array[i];

    printf("Notification: id:[%s] | score: [%f] | probability: [%f] | "
           "selected: [%d] | m_plus: [%f] | m_minus: [%f] | timestamp: [%ld]\n",
           notification->id, notification->score, notification->probability,
           notification->selected, notification->m_plus, notification->m_minus,
           notification->timestamp);
  }
  printf("\n--- Notification List ---\n");
}
