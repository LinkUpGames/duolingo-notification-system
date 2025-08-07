import setuptools

module = setuptools.Extension(
    "scorer",
    sources=[
        "module/python.c",
        "module/algorithm.c",
        "module/decision.c",
        "module/notification.c",
        "module/map.c",
    ],
    libraries=["m"],
)

setuptools.setup(
    name="scorer",
    version="0.1",
    packages=["scorer"],
    package_data={"scorer": ["__init__.pyi", "py.typed"]},
    ext_modules=[module],
    zip_safe=False,
)
