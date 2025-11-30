export interface User {
  id: number;
  username: string;
  role: 'admin' | 'technician' | 'warehouse';
  createdAt: string;
  updatedAt: string;
  deletedAt?: string | null;
}

export interface UserInput {
  username: string;
  password?: string;
  role?: 'admin' | 'technician' | 'warehouse';
  fileId?: number; // nếu dùng file upload
}
