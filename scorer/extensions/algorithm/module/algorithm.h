#ifndef ALGORITHM_H
#define ALGORITHM_H

#include "decision.h"
#include "notification.h"
#include <Python.h>
#include <sys/time.h>
/*
 * Given all the decisions, compute the scores for every round using the
 * decisions array
 * @param decisions The decision array
 * @param notification The notification map
 */
void compute_scores(DecisionArray *decisions, NotificationArray *notifications);

/**
 * Get current time in milliseconds
 */
long long current_time_millis();

#endif
