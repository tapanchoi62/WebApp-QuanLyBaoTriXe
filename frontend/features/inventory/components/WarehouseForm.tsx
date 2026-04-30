'use client';

import React, { useEffect } from 'react';
import { useForm } from 'react-hook-form';
import { Button } from '@/components/ui/button';
import { useHeader } from '@/contexts/HeaderContext';
import http from '@/lib/axios';
import { FormInput } from '@/components/FormInput';
import { FormHookProvider } from '@/components/FormHookProvider';
import { toast } from 'sonner';
import { useRouter } from 'next/navigation';

interface WarehouseFormProps {
  id: string | undefined;
}

interface WarehouseFormValues {
  name: string;
  location: string;
}

export default function WarehouseForm({ id }: WarehouseFormProps) {
  const { setTitle } = useHeader();
  const router = useRouter();

  const defaultValues: WarehouseFormValues = {
    name: '',
    location: '',
  };

  const form = useForm<WarehouseFormValues>({
    defaultValues,
  });

  const {
    formState: { isSubmitting },
    reset,
  } = form;

  const fetchData = async (warehouseId: string) => {
    try {
      const response = await http.get(`/api/warehouses/${warehouseId}`);
      const warehouseData = response.data.data;
      reset({
        name: warehouseData.name,
        location: warehouseData.location,
      });
    } catch (error) {
      console.error('Fetch warehouse failed:', error);
      toast.error('Failed to load warehouse data');
    }
  };

  useEffect(() => {
    setTitle(id ? 'Edit Warehouse' : 'Create Warehouse');
    if (id) {
      fetchData(id);
    }
  }, [id, reset]);

  const onSubmit = async (data: WarehouseFormValues) => {
    try {
      if (id) {
        await http.put(`/api/warehouses/${id}`, data);
      } else {
        await http.post('/api/warehouses', data);
      }
      toast.success('Saved successfully!');
      router.push('/warehouses');
    } catch (error) {
      console.error('Save warehouse failed:', error);
      toast.error('Save failed');
    }
  };

  return (
    <FormHookProvider form={form} onSubmit={onSubmit}>
      <div className="flex flex-row items-center w-full gap-4">
        <div className="w-1/2">
          <FormInput name="name" label="Name" placeholder="Warehouse name" />
        </div>
        <div className="w-1/2">
          <FormInput name="location" label="Location" placeholder="Warehouse location" />
        </div>
      </div>

      <Button type="submit" disabled={isSubmitting}>
        {id ? 'Update Warehouse' : 'Create Warehouse'}
      </Button>
    </FormHookProvider>
  );
}
