import setuptools

module = setuptools.Extension("module", sources=["module/algorithm.c"])

setuptools.setup(
    name="module",
    version="0.1",
    ext_modules=[module],
    packages=["module"],
    package_data={"module": ["module.pyi", "algorithm.pyi", "py.typed"]},
    zip_safe=False,
)
