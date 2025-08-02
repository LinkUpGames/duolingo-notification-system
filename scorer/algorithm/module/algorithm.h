// An arm that corresponds to a notification to be there
typedef struct Arm {
  // The name of the arm
  char *name;

  // The number of times that the arm was selected
  float count;

  // The reward [0,1] for the arm
  float reward;

  // The time this arm was last sent
  float last_sent;
} Arm;

/**
 * Select the arm given an array of arms with prior information to choose from
 * @param arms The array of arms
 * @param length the length of the array
 * @param alpha The alpha value that is used to know how accurate the alpha
 * value is
 * @param beta The beta value that is used to know how accurate previous
 * counts are
 */
char *select_arm(Arm **arms, int length, float alpha, float beta);
