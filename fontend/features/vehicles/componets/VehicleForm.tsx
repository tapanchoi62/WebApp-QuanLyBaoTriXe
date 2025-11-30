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

interface VehicleFormProps {
  id: string | undefined;
}

interface VehicleFormValues {
  plate_number: string;
  year: number;
  model_vehicle: string;
  note?: string;
}

export default function VehicleForm({ id }: VehicleFormProps) {
  const { setTitle } = useHeader();

  const [, setData] = React.useState<VehicleFormValues | null>(null);
  const defaultValues = {
    plate_number: '',
    year: new Date().getFullYear(),
    model_vehicle: '',
    note: '',
  };
  const form = useForm<VehicleFormValues>({
    defaultValues: defaultValues,
  });

  const {
    formState: { isSubmitting },
    reset,
    control,
  } = form;

  const fetchData = async (vehicleId: string | null) => {
    try {
      const response = (await http.get(`/vehicles/${vehicleId}`)) as any;
      const vehicleData = response.data;
      setData(vehicleData.data);
      reset({
        plate_number: vehicleData.data.plate_number,
        year: vehicleData.data.year?.toString(),
        model_vehicle: vehicleData.data.model_vehicle,
        note: vehicleData.data.note,
      });
    } catch (error) {
      console.error('Fetch vehicle failed:', error);
      toast('Failed to load vehicle data');
    }
  };

  useEffect(() => {
    setTitle(id ? 'Edit Vehicle' : 'Create Vehicle');
    if (id) {
      fetchData(id!);
    }
  }, [id, reset]);

  const onSubmit = async (data: VehicleFormValues) => {
    try {
      if (id) {
        // Update vehicle
        await http.put(`/vehicles/${id}`, data);
      } else {
        // Create new vehicle
        await http.post('/vehicles', {
          ...data,
          year: Number(data.year),
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
          <FormInput name="plate_number" label="Plate Number" placeholder="ABC-123" />
        </div>
        <div className="w-1/2">
          <FormSelect
            name="year"
            control={control}
            options={[
              { label: '2026', value: '2026' },
              { label: '2025', value: '2025' },
              { label: '2024', value: '2024' },
              { label: '2023', value: '2023' },
              { label: '2022', value: '2022' },
              { label: '2021', value: '2021' },
              { label: '2020', value: '2020' },
              { label: '2019', value: '2019' },
              { label: '2018', value: '2018' },
            ]}
            placeholder="Select vehicle type"
            isRequired
            label="Year"
          />
        </div>
      </div>
      <div className="flex flex-row items-center w-full gap-4">
        <div className="w-1/2">
          <FormInput name="model_vehicle" label="Model" placeholder="Enter model" />
        </div>
        <div className="w-1/2">
          <FormInput name="note" label="Note" placeholder="Optional note" />
        </div>
      </div>

      <Button type="submit" disabled={isSubmitting}>
        {id ? 'Update Vehicle' : 'Create Vehicle'}
      </Button>
    </FormHookProvider>
  );
}
