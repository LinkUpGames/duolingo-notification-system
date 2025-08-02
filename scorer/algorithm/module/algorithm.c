#include "algorithm.h"
#include "date.h"
#include <math.h>
#include <stdlib.h>
#include <string.h>

/**
 * Get the softmax values for the arms based on the temperature
 * NOTE: This returns a pointer that must be FREED!
 * @param arms The array of arms
 * @param length The length of the array
 * @param beta The temperature chosen
 */
float *softmax_probabilities(Arm **arms, int length, float beta) {
  float *values;
  float sum_exps = 0;
  float max_reward = 0;

  // Calculate the max score
  for (int i = 0; i < length; i++) {
    Arm *arm = arms[i];
    max_reward = arm->reward > max_reward ? arm->reward : max_reward;
  }

  // Calculate the exponential function
  values = (float *)calloc(length, sizeof(float));
  if (values == NULL) {
    return NULL;
  }

  for (int i = 0; i < length; i++) {
    Arm *arm = arms[i];
    float reward = arm->reward;

    float exp_p = exp((reward - max_reward) / beta);
    values[i] = exp_p;
    sum_exps += exp_p;
  }

  // Calculate softmax probability
  float *probabilities = (float *)calloc(length, sizeof(float));
  if (probabilities == NULL) {
    {
      return NULL;
    }

    for (int i = 0; i < length; i++) {
      float e = values[i];

      probabilities[i] = e / sum_exps;
    }
  }

  return probabilities;
}

void update_scores(Arm **arms, int length, char *chosen_arm, int selected,
                   float alpha, float beta) {
  // Compute Softwmax probabilities
  float *probabilities = softmax_probabilities(arms, length, beta);

  // Compute baseline
  float baseline = 0.0;
  for (int i = 0; i < length; i++) {
    baseline += probabilities[i] * arms[i]->reward;
  }

  // Update the Scores of each arm
  for (int i = 0; i < length; i++) {
    Arm *arm = arms[i];
    if (strcmp(arm->name, chosen_arm) == 0) {
      // Update based on previous selection
      arm->reward = alpha * (selected - baseline) * (1 - probabilities[i]);
    } else {
      arm->reward = alpha * (selected - baseline) * probabilities[i];
    }
  }
}
