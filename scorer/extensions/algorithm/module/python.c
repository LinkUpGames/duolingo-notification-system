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

  // Parse the list
  Decision **decision = parse_python_list(list);

  return NULL;
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
