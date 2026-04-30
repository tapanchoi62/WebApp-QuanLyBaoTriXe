'use client';

import React, { useEffect, useState } from 'react';
import { useForm } from 'react-hook-form';
import { Button } from '@/components/ui/button';
import { useHeader } from '@/contexts/HeaderContext';
import http from '@/lib/axios';
import { FormSelect } from '@/components/FormSelect';
import { FormInput } from '@/components/FormInput';
import { FormHookProvider } from '@/components/FormHookProvider';
import { toast } from 'sonner';
import { useRouter } from 'next/navigation';
import { Warehouse } from '../models/warehouse';

interface Item {
  ID: number;
  name: string;
}

interface Supplier {
  ID: number;
  name: string;
}

interface StockTransactionFormValues {
  warehouse_id: string;
  item_id: string;
  type: string;
  quantity: string;
  supplier_id: string;
  note: string;
}

export default function StockTransactionForm() {
  const { setTitle } = useHeader();
  const router = useRouter();
  const [warehouses, setWarehouses] = useState<Warehouse[]>([]);
  const [items, setItems] = useState<Item[]>([]);
  const [suppliers, setSuppliers] = useState<Supplier[]>([]);

  const defaultValues: StockTransactionFormValues = {
    warehouse_id: '',
    item_id: '',
    type: '',
    quantity: '',
    supplier_id: '',
    note: '',
  };

  const form = useForm<StockTransactionFormValues>({
    defaultValues,
  });

  const {
    formState: { isSubmitting },
    control,
  } = form;

  useEffect(() => {
    setTitle('Stock Transaction');

    const fetchOptions = async () => {
      try {
        const [warehousesRes, itemsRes, suppliersRes] = await Promise.all([
          http.get('/api/warehouses'),
          http.get('/api/items'),
          http.get('/api/suppliers'),
        ]);
        setWarehouses(warehousesRes.data.data);
        setItems(itemsRes.data.data);
        setSuppliers(suppliersRes.data.data);
      } catch (error) {
        console.error('Fetch options failed:', error);
      }
    };

    fetchOptions();
  }, [setTitle]);

  const onSubmit = async (data: StockTransactionFormValues) => {
    try {
      const payload = {
        warehouse_id: parseInt(data.warehouse_id, 10),
        item_id: parseInt(data.item_id, 10),
        type: data.type,
        quantity: parseFloat(data.quantity),
        supplier_id: data.supplier_id ? parseInt(data.supplier_id, 10) : null,
        note: data.note,
      };
      await http.post('/api/stocks/transaction', payload);
      toast.success('Transaction saved successfully!');
      router.push('/inventory');
    } catch (error) {
      console.error('Save transaction failed:', error);
      toast.error('Save failed');
    }
  };

  const warehouseOptions = warehouses.map((w) => ({
    label: w.name,
    value: String(w.ID),
  }));

  const itemOptions = items.map((i) => ({
    label: i.name,
    value: String(i.ID),
  }));

  const typeOptions = [
    { label: 'Nhập kho (IN)', value: 'IN' },
    { label: 'Xuất kho (OUT)', value: 'OUT' },
    { label: 'Điều chỉnh (ADJUST)', value: 'ADJUST' },
  ];

  const supplierOptions = suppliers.map((s) => ({
    label: s.name,
    value: String(s.ID),
  }));

  return (
    <FormHookProvider form={form} onSubmit={onSubmit}>
      <div className="flex flex-row items-start w-full gap-4">
        <div className="w-1/2">
          <FormSelect
            name="warehouse_id"
            control={control}
            options={warehouseOptions}
            placeholder="Select warehouse"
            isRequired
            label="Warehouse"
          />
        </div>
        <div className="w-1/2">
          <FormSelect
            name="item_id"
            control={control}
            options={itemOptions}
            placeholder="Select item"
            isRequired
            label="Item"
          />
        </div>
      </div>

      <div className="flex flex-row items-start w-full gap-4">
        <div className="w-1/2">
          <FormSelect
            name="type"
            control={control}
            options={typeOptions}
            placeholder="Select transaction type"
            isRequired
            label="Type"
          />
        </div>
        <div className="w-1/2">
          <FormInput name="quantity" label="Quantity" placeholder="Enter quantity" type="number" />
        </div>
      </div>

      <div className="flex flex-row items-start w-full gap-4">
        <div className="w-1/2">
          <FormSelect
            name="supplier_id"
            control={control}
            options={supplierOptions}
            placeholder="Select supplier"
            label="Supplier (optional)"
          />
        </div>
        <div className="w-1/2">
          <FormInput name="note" label="Note" placeholder="Optional note" />
        </div>
      </div>

      <Button type="submit" disabled={isSubmitting}>
        Submit Transaction
      </Button>
    </FormHookProvider>
  );
}
