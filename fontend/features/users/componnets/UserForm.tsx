'use client';

import { use, useEffect, useState } from 'react';
import { useForm } from 'react-hook-form';
import { Button } from '@/components/ui/button';
import { useHeader } from '@/contexts/HeaderContext';
import http from '@/lib/axios';
import { FormSelect } from '@/components/FormSelect';
import { FormInput } from '@/components/FormInput';
import { FormHookProvider } from '@/components/FormHookProvider';
import { toast } from 'sonner';

interface UserFormProps {
  id?: string; // nếu có id => edit, ngược lại => create
}

interface UserFormValues {
  username: string;
  role: string;
  password?: string;
}

interface Role {
  id: number;
  name: string;
}

export default function UserForm({ id }: UserFormProps) {
  const { setTitle } = useHeader();
  const [roles, setRoles] = useState<Role[]>([]);
  const defaultValues: UserFormValues = {
    username: '',
    role: '',
    password: '',
  };

  const form = useForm<UserFormValues>({
    defaultValues,
  });

  const {
    control,
    reset,
    formState: { isSubmitting },
  } = form;

  // Lấy danh sách role từ backend
  const fetchRoles = async () => {
    try {
      const res = await http.get('/roles');
      setRoles(res.data.data.map((r: any) => ({ id: r.ID, name: r.Name })));
    } catch (error) {
      console.error('Failed to load roles:', error);
      toast.error('Failed to load roles');
    }
  };

  // Lấy dữ liệu user nếu đang edit
  const fetchUser = async (userId: string) => {
    try {
      const res = await http.get(`/users/${userId}`);
      const user = res.data.data;
      const roleId = roles.find((r) => r.id === user.role?.ID)?.id.toString() || '';

      reset({
        username: user.username,
        role: roleId,
      });
    } catch (error) {
      console.error('Failed to load user:', error);
      toast.error('Failed to load user data');
    }
  };

  useEffect(() => {
    setTitle(id ? 'Edit User' : 'Create User');
    fetchRoles();
  }, []);

  useEffect(() => {
    if (id) {
      fetchUser(id);
    }
  }, [roles]);

  const onSubmit = async (data: UserFormValues) => {
    try {
      if (id) {
        // Update user
        await http.put(`/users/${id}`, {
          username: data.username,
          role_id: parseInt(data.role),
        });
      } else {
        // Create user
        await http.post('/users', {
          username: data.username,
          password: data.password,
          role_id: parseInt(data.role),
        });
      }
      toast.success('Saved successfully!');
    } catch (error) {
      console.error('Save failed:', error);
      toast.error('Save failed');
    }
  };

  return (
    <FormHookProvider form={form} onSubmit={onSubmit}>
      <div className="space-y-4 max-w-md">
        <FormInput
          name="username"
          label="Username"
          placeholder="Input username"
          readOnly={!!id} // edit không cho sửa username
        />
        {!id && (
          <FormInput
            name="password"
            label="Password"
            placeholder="Input password"
            readOnly={!!id} // edit không cho sửa username
          />
        )}

        <FormSelect
          name="role"
          control={control}
          options={roles.map((r) => ({ label: r.name, value: r.id.toString() }))}
          placeholder="Select role"
          label="Role"
          isRequired
        />

        <Button type="submit" disabled={isSubmitting}>
          {id ? 'Update User' : 'Create User'}
        </Button>
      </div>
    </FormHookProvider>
  );
}
