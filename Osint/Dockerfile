FROM debian

# Setting up bash as shell
SHELL ["/bin/bash", "-c"]

# Installing needed packages and creating needed directories
RUN apt update && apt install -y python3 python3.11-venv pip firefox-esr && mkdir app && mkdir env

# Getting required files
WORKDIR /app
COPY osint.py /app/osint.py
COPY requirements.txt /app/requirements.txt

# Setting up the python virtual environement
RUN python3 -m venv /app/env && source /app/env/bin/activate && pip install -r requirements.txt

# Addind the script to launch
RUN echo "source /app/env/bin/activate && python3 osint.py" > launch.sh

ENTRYPOINT ["/bin/bash",  "launch.sh"]