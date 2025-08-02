#include "date.h"

// An arm that corresponds to a notification to be there
typedef struct Arm {
  // The name of the arm
  char *name;

  // The number of times that the arm was selected
  float count;

  // The reward [0,1] for the arm
  float reward;

  // The time this arm was last sent as a timestamp
  timestamp last_sent;
} Arm;

/**
 * Update the scores of each arm given the prior chosen arm
 * @param arms Array with each arm
 * @param length The length of the arms array
 * @param chosen_arm The name of the arm that was chosen by algorithm previously
 * @param selected Whether the arm by algorithm was selected by the user (1 or
 * 0);
 * @param alpha The learning rate
 * @param beta The temperature
 */
void update_scores(Arm **arms, int length, char *chosen_arm, int selected,
                   float alpha, float beta);
