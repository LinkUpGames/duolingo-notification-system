#include "algorithm.h"
#include "decision.h"
#include "notification.h"
#include <sys/time.h>

void compute_scores(DecisionArray *decisions,
                    NotificationArray *notifications) {

  // Decay rate hyperparameter
  const double lambda_decay = 0.001;
  const double beta = 1.0;

  const long long now = current_time_millis();

  // Map from notification ID -> Notification*
  hashmap *map = create_notification_map_from_list(notifications);
  if (map == NULL) {
    PyErr_NoMemory();
    return;
  }

  // Loop through each decision
  for (size_t i = 0; i < decisions->length; i++) {
    Decision *decision = decisions->array[i];
    const char *selected_id = decision->notification_id;
    const int selected = decision->selected;

    // Update the values of all notifications
    for (size_t j = 0; j < decision->probabilities->length; j++) {
      // probability
      Notification *p = decision->probabilities->array[j];
      double probability = p->probability;

      uintptr_t map_value;
      hashmap_get(map, p->id, strlen(p->id), &map_value);

      // The selected notification
      if (map_value) {
        Notification *notification = (Notification *)(map_value);

        double elapsed =
            (notification->timestamp > 0) ? now - notification->timestamp : 0.0;
        double decay_factor = exp(-lambda_decay * elapsed);
        notification->m_plus *= decay_factor;
        notification->m_minus *= decay_factor;

        if (strcmp(notification->id, selected_id) == 0) {
          if (selected == 1) {
            notification->m_plus += 1.0;
          } else {
            notification->m_minus += 1.0;
          }
        } else {
          notification->m_minus += probability;
        }
      }
    }
  }

  // update Scores
  for (size_t j = 0; j < notifications->length; j++) {
    Notification *notification = notifications->array[j];

    notification->score =
        1.0 /
        (1.0 + exp(-(notification->m_plus - notification->m_minus) / beta));
  }

  // Free hashmap
  hashmap_free(map);
}

long long current_time_millis() {
  struct timeval tv;
  gettimeofday(&tv, NULL);
  return (long long)tv.tv_sec * 1000 + tv.tv_usec / 1000;
}
