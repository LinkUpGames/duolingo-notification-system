#define PY_SSIZE_T_CLEAN
#include <Python.h>
#include <stdio.h>

static PyObject *say_hello(PyObject *self, PyObject *args) {
  const char *name;

  if (!PyArg_ParseTuple(args, "s", &name))
    return NULL;

  printf("Hello, %s\n", name);

  Py_RETURN_NONE;
}

static PyMethodDef MyMethods[] = {
    {"say_hello", say_hello, METH_VARARGS, "Print hello message."},
    {NULL, NULL, 0, NULL},
};

static struct PyModuleDef mymodule = {
    PyModuleDef_HEAD_INIT, "module", NULL, -1, MyMethods,
};

PyMODINIT_FUNC PyInit_my_module(void) { return PyModule_Create(&mymodule); }
