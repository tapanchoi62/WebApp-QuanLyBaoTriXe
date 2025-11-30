'use client';

import React, { useEffect } from 'react';
import { useForm } from 'react-hook-form';
import { Button } from '@/components/ui/button';
import { useHeader } from '@/contexts/HeaderContext';
import http from '@/lib/axios';
import { FormInput } from '@/components/FormInput';
import { FormHookProvider } from '@/components/FormHookProvider';
import { toast } from 'sonner';

interface ItemFormProps {
  id: string | undefined;
}

interface ItemFormValues {
  category: string;
  unit: string;
  name: string;
}

export default function VehicleForm({ id }: ItemFormProps) {
  const { setTitle } = useHeader();

  const [, setData] = React.useState<ItemFormValues | null>(null);
  const defaultValues = {
    category: '',
    unit: '',
    name: '',
  };
  const form = useForm<ItemFormValues>({
    defaultValues: defaultValues,
  });

  const {
    formState: { isSubmitting },
    reset,
  } = form;

  const fetchData = async (itemId: string | null) => {
    try {
      const response = (await http.get(`/items/${itemId}`)) as any;
      const itemData = response.data;
      setData(itemData.data);
      reset({
        name: itemData.data.name,
        unit: itemData.data.unit,
        category: itemData.data.category,
      });
    } catch (error) {
      console.error('Fetch item failed:', error);
      toast('Failed to load item data');
    }
  };

  useEffect(() => {
    setTitle(id ? 'Edit item' : 'Create item');
    if (id) {
      fetchData(id!);
    }
  }, [id, reset]);

  const onSubmit = async (data: ItemFormValues) => {
    try {
      if (id) {
        // Update vehicle
        await http.put(`/items/${id}`, data);
      } else {
        // Create new vehicle
        await http.post('/items', data);
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
          <FormInput name="name" label="Name" placeholder="Enter name" />
        </div>
        <div className="w-1/2">
          <FormInput name="unit" label="Unit" placeholder="Enter unit" />
        </div>
        <div className="w-1/2">
          <FormInput name="category" label="Category" placeholder="Enter category" />
        </div>
      </div>

      <Button type="submit" disabled={isSubmitting}>
        {id ? 'Update Vehicle' : 'Create Vehicle'}
      </Button>
    </FormHookProvider>
  );
}
