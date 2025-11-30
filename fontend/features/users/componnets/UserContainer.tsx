'use client';

import { useEffect, useState } from 'react';
import { Button } from '@/components/ui/button';
import {
  Table,
  TableHeader,
  TableRow,
  TableHead,
  TableBody,
  TableCell,
} from '@/components/ui/table';
import { usePermission } from '@/hooks/use-permission';
import http from '@/lib/axios';
import { useRouter } from 'next/navigation';

interface Permission {
  ID: number;
  Name: string;
}

interface Role {
  ID: number;
  Name: string;
  Permissions: Permission[];
}

interface User {
  ID: number;
  Username: string;
  Role: Role;
}

export default function PageUser() {
  const [users, setUsers] = useState<User[]>([]);
  const router = useRouter();
  // Quyền
  const canCreate = usePermission('CREATE_USER');
  const canEdit = usePermission('EDIT_USER');
  const canDelete = usePermission('DELETE_USER');
  const SYS_ADMIN = usePermission('SYS_ADMIN');

  const loadUsers = async () => {
    try {
      const res = await http.get('/users');
      setUsers(res.data.data);
    } catch (err) {
      console.error(err);
    }
  };

  useEffect(() => {
    loadUsers();
  }, []);

  const deleteUser = async (id: number) => {
    if (!confirm('Bạn có chắc muốn xóa user này?')) return;
    await http.delete(`/users/${id}`);
    loadUsers();
  };

  return (
    <div className="p-6 space-y-4">
      <div className="flex justify-between items-center">
        <h1 className="text-xl font-bold">Quản lý User</h1>
        {(canCreate || SYS_ADMIN) && (
          <Button onClick={() => router.push('/manager/users/create')}>+ Tạo User</Button>
        )}
      </div>

      <Table>
        <TableHeader>
          <TableRow>
            <TableHead>ID</TableHead>
            <TableHead>Username</TableHead>
            <TableHead>Role</TableHead>
            <TableHead>Permissions</TableHead>
            <TableHead className="text-right">Actions</TableHead>
          </TableRow>
        </TableHeader>

        <TableBody>
          {users.map((u) => (
            <TableRow key={u.ID}>
              <TableCell>{u.ID}</TableCell>
              <TableCell>{u.Username}</TableCell>
              <TableCell>{u.Role?.Name}</TableCell>
              <TableCell>
                {u.Role?.Permissions?.map((p) => (
                  <span
                    key={p.ID}
                    className="inline-block bg-gray-200 text-gray-800 px-2 py-1 rounded mr-1 text-xs"
                  >
                    {p.Name}
                  </span>
                ))}
              </TableCell>
              <TableCell className="text-right space-x-2">
                {(canEdit || SYS_ADMIN) && (
                  <Button
                    size="sm"
                    variant="outline"
                    onClick={() => router.push('/manager/users/' + u.ID)}
                  >
                    Edit
                  </Button>
                )}
                {(canDelete || SYS_ADMIN) && (
                  <Button size="sm" variant="destructive" onClick={() => deleteUser(u.ID)}>
                    Delete
                  </Button>
                )}
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </div>
  );
}
