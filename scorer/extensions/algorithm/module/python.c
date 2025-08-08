#define PY_SSIZE_T_CLEAN
#include "algorithm.h"
#include <Python.h>

/**
 * Wrapper for the python method that computes the scores given a list of
 * decision logs
 * @param self The calling object
 * @param args The argument list
 */
static PyObject *compute_scores_method(PyObject *self, PyObject *args) {
  PyObject *list;

  // Parse arguments
  if (!PyArg_ParseTuple(args, "O!", &PyList_Type, &list)) {
    PyErr_SetString(PyExc_TypeError, "Expected: list[Decision]");

    // Return null
    Py_RETURN_NONE;
  }

  return NULL;
}

static PyObject *update_scores_method(PyObject *self, PyObject *args) {
  PyObject *input_list;
  const char *chosen_arm;
  int selected;
  float alpha;
  float temperature;

  // Check for the right input
  if (!PyArg_ParseTuple(args, "O!siff", &PyList_Type, &input_list, &chosen_arm,
                        &selected, &alpha, &temperature)) {
    PyErr_SetString(PyExc_TypeError, "Expected: list, str, int, float, float");

    return NULL;
  }

  Py_ssize_t list_size = PyList_Size(input_list);

  // Parse the input
  Arm **arms = parse_python_input(input_list, list_size);

  // Update The scores
  update_scores(arms, list_size, chosen_arm, selected, alpha, temperature);

  PyObject *list = return_arms(arms, list_size);

  // Free memory
  free_arm_list(arms, list_size);

  return list;
}

static PyMethodDef scorer_methods[] = {
    {"update_scores", update_scores_method, METH_VARARGS,
     "Print hello message."},
    {"compute_scores", compute_scores_method, METH_VARARGS,
     "Compute the scores for the notification given a list of decisions"},
    {NULL, NULL, 0, NULL},
};

static struct PyModuleDef spam_module = {
    PyModuleDef_HEAD_INIT, "scorer", NULL, -1, scorer_methods,
};

PyMODINIT_FUNC PyInit_scorer(void) { return PyModule_Create(&spam_module); }
