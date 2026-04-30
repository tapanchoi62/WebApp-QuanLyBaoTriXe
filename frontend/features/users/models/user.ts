export interface User {
  ID: number;
  Username: string;
  Role: {
    ID: number;
    Name: string;
    Permissions: { ID: number; Name: string }[];
  };
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt: string | null;
}

export interface UserInput {
  username: string;
  password?: string;
  role_id?: number;
}
