export interface User {
  id: string
  full_name: string
  email: string
  role_id?: number
}

export interface UpdateUserBody {
  full_name?: string
  email?: string
  password?: string
}
