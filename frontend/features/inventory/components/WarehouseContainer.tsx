'use client';

import { useEffect, useState } from 'react';
import { ColumnDef } from '@tanstack/react-table';
import { DataTable } from '@/components/DataTable';
import http from '@/lib/axios';
import { useHeader } from '@/contexts/HeaderContext';
import { Warehouse } from '../models/warehouse';
import { useRouter } from 'next/navigation';
import { Button } from '@/components/ui/button';
import { Eye, Trash2 } from 'lucide-react';
import { toast } from 'sonner';
import { ConfirmButton } from '@/components/ConfirmButton';

export default function WarehouseContainer() {
  const [data, setData] = useState<Warehouse[]>([]);
  const [page, setPage] = useState(1);
  const pageSize = 20;
  const [total, setTotal] = useState(0);
  const [search, setSearch] = useState('');
  const router = useRouter();
  const { setTitle } = useHeader();

  useEffect(() => {
    setTitle('Warehouse Management');
  }, [setTitle]);

  const columns: ColumnDef<Warehouse>[] = [
    {
      header: 'ID',
      accessorKey: 'ID',
    },
    {
      header: 'Name',
      accessorKey: 'name',
    },
    {
      header: 'Location',
      accessorKey: 'location',
    },
    {
      header: 'Created At',
      accessorKey: 'CreatedAt',
      cell: (info) => new Date(info.getValue() as string).toLocaleString(),
    },
    {
      header: 'Actions',
      cell: (info) => {
        const row = info.row.original;
        return (
          <div className="flex flex-row gap-2">
            <Button
              onClick={() => router.push(`/warehouses/${row.ID}`)}
              variant={'outline'}
              className="rounded-2xl"
            >
              <Eye />
            </Button>
            <ConfirmButton
              variant="destructive"
              dialogTitle="Delete Warehouse"
              dialogDescription="Are you sure you want to delete this warehouse? This action cannot be undone."
              confirmText="Delete"
              cancelText="Cancel"
              onConfirm={() => removeWarehouse(row.ID)}
            >
              <Trash2 />
            </ConfirmButton>
          </div>
        );
      },
    },
  ];

  const removeWarehouse = async (id: number) => {
    try {
      await http.delete(`/api/warehouses/${id}`);
      toast.success('Warehouse deleted successfully');
      fetchData();
    } catch (error) {
      console.error('Delete warehouse failed:', error);
      toast.error('Failed to delete warehouse');
    }
  };

  const fetchData = async () => {
    const res = await http.get('/api/warehouses');
    setData(res.data.data);
    setTotal(res.data.data.length);
  };

  useEffect(() => {
    const timeout = setTimeout(() => {
      fetchData();
    }, 300);

    return () => clearTimeout(timeout);
  }, [page, search]);

  const onCreate = () => {
    router.push('/warehouses/create');
  };

  return (
    <div className="p-4">
      <DataTable
        onSearchChange={setSearch}
        columns={columns}
        data={data}
        page={page}
        pageSize={pageSize}
        total={total}
        onPageChange={setPage}
        onCreate={onCreate}
        search={search}
      />
    </div>
  );
}
