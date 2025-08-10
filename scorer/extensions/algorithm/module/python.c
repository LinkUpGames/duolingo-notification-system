#define PY_SSIZE_T_CLEAN
#include "algorithm.h"
#include "decision.h"
#include "notification.h"
#include <Python.h>

/**
 * Wrapper for the python method that computes the scores given a list of
 * decision logs
 * @param self The calling object
 * @param args The argument list
 */
static PyObject *compute_scores_method(PyObject *self, PyObject *args) {
  PyObject *decision_list_obj;
  PyObject *notification_list_obj;

  // Parse arguments
  if (!PyArg_ParseTuple(args, "O!O!", &PyList_Type, &decision_list_obj,
                        &PyList_Type, &notification_list_obj)) {
    PyErr_SetString(PyExc_TypeError,
                    "Expected: list[Decision], list[Notifications]");

    // Return null
    Py_RETURN_NONE;
  }

  // Parse the list
  DecisionArray *decisions = parse_python_decision_list(decision_list_obj);
  NotificationArray *notifications =
      parse_python_notification_list(notification_list_obj);

  // Compute Scores
  compute_scores(decisions, notifications);

  free_decision_list(decisions);
  free_notification_list(notifications);

  Py_RETURN_NONE;
}

static PyMethodDef scorer_methods[] = {
    {"compute_scores", compute_scores_method, METH_VARARGS,
     "Compute the scores for the notification given a list of decisions"},
    {NULL, NULL, 0, NULL},
};

static struct PyModuleDef spam_module = {
    PyModuleDef_HEAD_INIT, "scorer", NULL, -1, scorer_methods,
};

PyMODINIT_FUNC PyInit_scorer(void) { return PyModule_Create(&spam_module); }
