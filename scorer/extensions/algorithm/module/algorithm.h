#ifndef ALGORITHM_H
#define ALGORITHM_H

#include "decision.h"
#include "notification.h"
#include <Python.h>

/*
 * Given all the decisions, compute the scores for every round using the
 * decisions array
 * @param decisions The decision array
 * @param length The lengthof the array
 */
NotificationArray *compute_scores(DecisionArray *decision, long length);

#endif
