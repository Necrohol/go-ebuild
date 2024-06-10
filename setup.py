try:
    from setuptools import setup
except ImportError:
    from distutils.core import setup


setup(
    name='go-ebuild',
    long_description=open('README.md').read(),
    long_description_content_type='text/markdown',
    packages=[
        'go_ebuild',
        'tests',
    ],
)
