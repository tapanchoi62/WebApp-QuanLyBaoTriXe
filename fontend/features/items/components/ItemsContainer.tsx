'use client';

import { useEffect, useState } from 'react';
import { ColumnDef } from '@tanstack/react-table';
import { DataTable } from '@/components/DataTable';
import http from '@/lib/axios';
import { useHeader } from '@/contexts/HeaderContext';
import { useRouter } from 'next/navigation';
import { Button } from '@/components/ui/button';
import { Eye, Trash2 } from 'lucide-react';
import { toast } from 'sonner';
import { ConfirmButton } from '@/components/ConfirmButton';
import { Item } from '../models/item';

export default function ItemsContainer() {
  const [data, setData] = useState<Item[]>([]);
  const [page, setPage] = useState(1);
  const pageSize = 20;
  const [total, setTotal] = useState(0);
  const [search, setSearch] = useState('');
  const router = useRouter();
  const { setTitle } = useHeader();

  useEffect(() => {
    setTitle('Item Management');
  }, [setTitle]);

  const columns: ColumnDef<Item>[] = [
    {
      header: 'ID',
      accessorKey: 'ID',
    },
    {
      header: 'name',
      accessorKey: 'name',
    },
    {
      header: 'category',
      accessorKey: 'category',
    },
    {
      header: 'unit',
      accessorKey: 'unit',
    },
    {
      header: 'Actions',
      cell: (info) => {
        const row = info.row.original;
        return (
          <div className="flex flex-row gap-2">
            <Button
              onClick={() => router.push(`/items/${row.ID}`)}
              variant={'outline'}
              className="rounded-2xl"
            >
              <Eye></Eye>
            </Button>
            <ConfirmButton
              variant="destructive"
              dialogTitle="Delete Item"
              dialogDescription="Are you sure you want to delete this Item? This action cannot be undone."
              confirmText="Delete"
              cancelText="Cancel"
              onConfirm={() => removeItem(row.ID)}
            >
              <Trash2></Trash2>
            </ConfirmButton>
          </div>
        );
      },
    },
  ];

  const removeItem = async (id: number) => {
    try {
      await http.delete(`/items/${id}`);
      toast('Item deleted successfully');
      fetchData();
    } catch (error) {
      console.error('Delete Item failed:', error);
      toast('Failed to delete Item');
    }
  };
  const fetchData = async () => {
    const res = await http.get('/items', {
      params: { page, pageSize, search },
    });

    setData(res.data.data);
    setTotal(res.data.pagination.total);
  };

  useEffect(() => {
    const timeout = setTimeout(() => {
      fetchData();
    }, 300); // 300ms debounce

    return () => clearTimeout(timeout);
  }, [page, search]);

  const onCreate = () => {
    router.push('/items/create');
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
