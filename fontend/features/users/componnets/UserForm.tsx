'use client';

import React, { useEffect } from 'react';
import { useForm } from 'react-hook-form';
import { Button } from '@/components/ui/button';
import { useHeader } from '@/contexts/HeaderContext';
import http from '@/lib/axios';
import { FormSelect } from '@/components/FormSelect';
import { FormInput } from '@/components/FormInput';
import { FormHookProvider } from '@/components/FormHookProvider';
import { toast } from 'sonner';

interface UserFormProps {
  id: string | undefined;
}

interface UserFormValues {
  username: string;
  role: string;
  fileId?: string;
}

export default function UserForm({ id }: UserFormProps) {
  const { setTitle } = useHeader();

  const [, setData] = React.useState<UserFormValues | null>(null);
  const defaultValues = {
    username: '',
    role: '',
    fileId: '',
  };
  const form = useForm<UserFormValues>({
    defaultValues: defaultValues,
  });

  const {
    formState: { isSubmitting },
    reset,
    control,
  } = form;

  const fetchData = async (userId: string | null) => {
    try {
      const response = (await http.get(`/users/${userId}`)) as any;
      const usersData = response.data;
      setData(usersData.data);
      reset({
        username: usersData.data.username,
        role: usersData.data.role?.toString(),
        fileId: usersData.data.fileId?.toString(),
      });
    } catch (error) {
      console.error('Fetch user failed:', error);
      toast('Failed to load user data');
    }
  };

  useEffect(() => {
    setTitle(id ? 'Edit user' : 'Create user');
    if (id) {
      fetchData(id!);
    }
  }, [id, reset]);

  const onSubmit = async (data: UserFormValues) => {
    try {
      if (id) {
        // Update vehicle
        await http.put(`/users/${id}`, data);
      } else {
        // Create new vehicle
        await http.post('/users', {
          ...data,
        });
      }
      toast.success('Saved successfully!');
    } catch (error) {
      console.error('Save vehicle failed:', error);
      toast.error('Save failed');
    }
  };

  return (
    <FormHookProvider form={form} onSubmit={onSubmit}>
      <div className="flex flex-row items-center w-full gap-4">
        <div className="w-1/2">
          <FormInput name="username" label="Username" placeholder="Input Username" />
        </div>
        <div className="w-1/2">
          <FormSelect
            name="role"
            control={control}
            options={[
              { label: 'technician', value: 'technician' },
              { label: 'admin', value: 'admin' },
              { label: 'warehouse', value: 'warehouse' },
            ]}
            placeholder="Select vehicle type"
            isRequired
            label="Year"
          />
        </div>
      </div>

      <Button type="submit" disabled={isSubmitting}>
        {id ? 'Update Vehicle' : 'Create Vehicle'}
      </Button>
    </FormHookProvider>
  );
}
