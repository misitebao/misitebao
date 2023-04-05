from setuptools import setup

setup(
    name='misitebao',
    version='1.0.2',
    keywords=["misitebao", "misi"],
    license="MIT License",
    author='Misite Bao',
    author_email='i@misitebao.com',
    url='https://github.com/misitebao/misitebao',
    description='Hi, I am Misite Bao.',
    entry_points={
        'console_scripts': [
            'misitebao=main:main',
            'misi=main:main'
        ]
    }
)
