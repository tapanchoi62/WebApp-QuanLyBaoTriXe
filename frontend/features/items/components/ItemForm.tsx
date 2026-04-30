'use client';

import React, { useEffect } from 'react';
import { useForm } from 'react-hook-form';
import { Button } from '@/components/ui/button';
import { useHeader } from '@/contexts/HeaderContext';
import http from '@/lib/axios';
import { FormInput } from '@/components/FormInput';
import { FormSelect } from '@/components/FormSelect';
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

const CATEGORY_OPTIONS = [
  { label: 'Tire', value: 'Tire' },
  { label: 'Battery', value: 'Battery' },
  { label: 'Oil', value: 'Oil' },
  { label: 'Filter', value: 'Filter' },
  { label: 'Brake', value: 'Brake' },
  { label: 'Lighting', value: 'Lighting' },
  { label: 'Belt', value: 'Belt' },
  { label: 'Cooling', value: 'Cooling' },
  { label: 'Fluid', value: 'Fluid' },
];
const UNIT_OPTIONS = [
  { label: 'cái', value: 'cái' },
  { label: 'bình', value: 'bình' },
  { label: 'lít', value: 'lít' },
  { label: 'bộ', value: 'bộ' },
  { label: 'sợi', value: 'sợi' },
];

export default function ItemForm({ id }: ItemFormProps) {
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
    control,
  } = form;

  const fetchData = async (itemId: string | null) => {
    try {
      const response = (await http.get(`/api/items/${itemId}`)) as any;
      const itemData = response.data;
      setData(itemData.data);
      reset({
        name: itemData.data.name,
        category: itemData.data.category,
        unit: itemData.data.unit,
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
        // Update item
        await http.put(`/api/items/${id}`, data);
      } else {
        // Create new item
        await http.post('/api/items', data);
      }
      toast.success('Saved successfully!');
    } catch (error) {
      console.error('Save item failed:', error);
      toast.error('Save failed');
    }
  };

  return (
    <FormHookProvider form={form} onSubmit={onSubmit}>
      <div className="flex flex-row items-center w-full gap-4">
        <div className="w-full">
          <FormInput name="name" label="Name" placeholder="Enter name" />
        </div>
      </div>
      <div className="flex flex-row items-center w-full gap-4">
        <div className="w-1/2">
          <FormSelect name="category" control={control} options={CATEGORY_OPTIONS} label="Category" isRequired />
        </div>
        <div className="w-1/2">
          <FormSelect name="unit" control={control} options={UNIT_OPTIONS} label="Unit" isRequired />
        </div>
      </div>

      <Button type="submit" disabled={isSubmitting}>
        {id ? 'Update Item' : 'Create Item'}
      </Button>
    </FormHookProvider>
  );
}
