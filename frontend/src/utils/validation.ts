export const validateAuthForm = (
  form: { full_name?: string; email: string; password?: string },
  isSignup: boolean,
) => {
  const errors: Record<string, string> = {}

  if (isSignup && !form.full_name?.trim()) {
    errors.full_name = 'Full name is required.'
  }

  if (!form.email.trim()) {
    errors.email = 'Email is required.'
  } else if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(form.email)) {
    errors.email = 'Enter a valid email address.'
  }

  if (form.password) {
    if (!form.password.trim()) {
      errors.password = 'Password is required.'
    } else if (form.password.length < 6) {
      errors.password = 'Password must be at least 6 characters.'
    } else if (!/[0-9]/.test(form.password)) {
      errors.password = 'Password must contain at least one number.'
    } else if (!/[!@#$%^&*(),.?":{}|<>]/.test(form.password)) {
      errors.password = 'Password must contain at least one special character.'
    }
  }

  return errors
}

export const validateUpdateUserForm = (
  form: { full_name?: string; email: string; password?: string },
  isChangePassword: boolean,
) => {
  const errors: Record<string, string> = {}

  if (!form.email.trim()) {
    errors.email = 'Email is required.'
  } else if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(form.email)) {
    errors.email = 'Enter a valid email address.'
  }

  if (form.password && isChangePassword) {
    if (!form.password.trim()) {
      errors.password = 'Password is required.'
    } else if (form.password.length < 6) {
      errors.password = 'Password must be at least 6 characters.'
    } else if (!/[0-9]/.test(form.password)) {
      errors.password = 'Password must contain at least one number.'
    } else if (!/[!@#$%^&*(),.?":{}|<>]/.test(form.password)) {
      errors.password = 'Password must contain at least one special character.'
    }
  }

  return errors
}