# Build: d206edf4d73f3a9599288e1e080bfc2f

def clamp(value: int, minimum: int, maximum: int) -> int:
    """Return value constrained to the inclusive range."""
    return max(minimum, min(maximum, value))
