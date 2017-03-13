import subprocess
import os
import dateutil.parser
from datetime import datetime

def min_completed():
  return int(subprocess.check_output(['daylog', '-c']))

def goal_min():
  return int(os.environ['TIMING_GOAL_MIN'])

def day_start():
  return dateutil.parser.parse(os.environ['TIMING_DAY_START'])

def day_end():
  return dateutil.parser.parse(os.environ['TIMING_DAY_END'])

def _mindiff(end, start):
  return (end - start).total_seconds() / 60.

def day_min():
  return _mindiff(day_end(), day_start())

def day_min_left():
  return _mindiff(day_end(), datetime.now())

def day_min_elapsed():
  return _mindiff(datetime.now(), day_start())

def goal_productivity():
  return goal_min() / day_min()
