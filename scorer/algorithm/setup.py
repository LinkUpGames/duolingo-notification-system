import setuptools

module = setuptools.Extension("my_module", sources=["algorithm.c"])

setuptools.setup(name="my_module", version="0.1", ext_modules=[module])
