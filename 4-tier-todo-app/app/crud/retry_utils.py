from functools import wraps
from retry import retry
from mysql.connector.errors import OperationalError

def retry_on_operational_error(func):
    @wraps(func)
    @retry(OperationalError, delay=1, backoff=2, max_delay=10, tries=5)
    def wrapper(*args, **kwargs):
        return func(*args, **kwargs)
    return wrapper
