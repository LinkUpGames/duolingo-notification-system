#define PY_SSIZE_T_CLEAN
#include "algorithm.h"
#include <Python.h>

static PyObject *update_scores_method(PyObject *self, PyObject *args) {
  PyObject *input_list;
  const char *chosen_arm;
  int selected;
  float alpha;
  float temperature;

  // Check for the right input
  if (!PyArg_ParseTuple(args, "O!sidd", &PyList_Type, &input_list, &chosen_arm,
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
    {NULL, NULL, 0, NULL},
};

static struct PyModuleDef spam_module = {
    PyModuleDef_HEAD_INIT, "scorer", NULL, -1, scorer_methods,
};

PyMODINIT_FUNC PyInit_scorer(void) { return PyModule_Create(&spam_module); }
