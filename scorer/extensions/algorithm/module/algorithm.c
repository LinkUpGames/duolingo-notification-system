#include "decision.h"
#include "notification.h"

void compute_scores(DecisionArray *decisions,
                    NotificationArray *notifications) {
  // Map
  hashmap *map = create_notification_map_from_list(notifications);
  if (map == NULL) {
    PyErr_NoMemory();

    return;
  }

  // Set the weighted accordingly
  for (size_t i = 0; i < decisions->length; i++) {
    Decision *decision = decisions->array[i];

    int was_selected = decision->selected;
    char *selected_notification_id = decision->notification_id;

    for (size_t j = 0; j < decision->probabilities->length; j++) {
      Notification *notification = decision->probabilities->array[j];

      double weight = 1 / notification->probability;
      double value = was_selected * weight;

      // Get the notification from the hasmap
      uintptr_t map_value;
      hashmap_get(map, notification->id, strlen(notification->id), &map_value);

      if (map_value) {
        if (strcmp(notification->id, selected_notification_id) ==
            0) { // This is the notification that was selected
          ((Notification *)map_value)->m_plus_count += value * weight;
          ((Notification *)map_value)->m_plus_sum += weight;
        } else {
          // Not selected, reweight score
          ((Notification *)map_value)->m_minus_count += value * weight;
          ((Notification *)map_value)->m_minus_sum += weight;
        }
      }
    }
  }

  // Calculate the m_plus and m_minus
  for (size_t i = 0; i < notifications->length; i++) {
    Notification *notification = notifications->array[i];

    float m_plus = 0;
    float m_minus = 1;

    m_plus = notification->m_plus_count > 0
                 ? notification->m_plus_sum / notification->m_plus_count
                 : m_plus;

    m_minus = notification->m_minus_count > 0
                  ? notification->m_minus_sum / notification->m_minus_count
                  : m_minus;

    // Scores
    notification->m_plus = m_plus;
    notification->m_minus = m_minus;
    notification->score = (m_plus / m_minus) / m_minus;
  }

  // Free this
  hashmap_free(map);
}
