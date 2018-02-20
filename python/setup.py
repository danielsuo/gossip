from setuptools import setup, find_packages

setup(
    name='gossip',
    version='0.0.1',
    description='Gossip protocols',
    url='https://github.com/danielsuo/gossip',
    author='Daniel Suo',
    author_email='dsuo@cs.princeton.edu',
    classifiers=[
        'Development Status :: 3 - Alpha',
        'License :: OSI Approved :: MIT License',
        'Programming Language :: Python :: 3.6'
    ],
    install_requires=[
        'click'
    ],
    packages=find_packages(),
    python_requires='>=3.6'
)
