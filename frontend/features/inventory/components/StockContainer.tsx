'use client';

import { useEffect, useState } from 'react';
import { ColumnDef } from '@tanstack/react-table';
import { DataTable } from '@/components/DataTable';
import http from '@/lib/axios';
import { useHeader } from '@/contexts/HeaderContext';
import { Stock } from '../models/stock';
import { Warehouse } from '../models/warehouse';
import { useRouter } from 'next/navigation';
import { Button } from '@/components/ui/button';
import { Eye } from 'lucide-react';

export default function StockContainer() {
  const [data, setData] = useState<Stock[]>([]);
  const [page, setPage] = useState(1);
  const pageSize = 20;
  const [total, setTotal] = useState(0);
  const [search, setSearch] = useState('');
  const [warehouses, setWarehouses] = useState<Warehouse[]>([]);
  const [selectedWarehouseId, setSelectedWarehouseId] = useState<string>('');
  const router = useRouter();
  const { setTitle } = useHeader();

  useEffect(() => {
    setTitle('Inventory (Stock)');
  }, [setTitle]);

  useEffect(() => {
    const fetchWarehouses = async () => {
      try {
        const res = await http.get('/api/warehouses');
        setWarehouses(res.data.data);
      } catch (error) {
        console.error('Fetch warehouses failed:', error);
      }
    };
    fetchWarehouses();
  }, []);

  const columns: ColumnDef<Stock>[] = [
    {
      header: 'ID',
      accessorKey: 'ID',
    },
    {
      header: 'Item Name',
      accessorKey: 'Item',
      cell: (info) => {
        const row = info.row.original;
        return row.Item?.name ?? '-';
      },
    },
    {
      header: 'Category',
      accessorKey: 'category',
      cell: (info) => {
        const row = info.row.original;
        return row.Item?.category ?? '-';
      },
    },
    {
      header: 'Unit',
      accessorKey: 'unit',
      cell: (info) => {
        const row = info.row.original;
        return row.Item?.unit ?? '-';
      },
    },
    {
      header: 'Warehouse',
      accessorKey: 'Warehouse',
      cell: (info) => {
        const row = info.row.original;
        return row.Warehouse?.name ?? '-';
      },
    },
    {
      header: 'Quantity',
      accessorKey: 'quantity',
    },
    {
      header: 'Actions',
      cell: (info) => {
        const row = info.row.original;
        return (
          <div className="flex flex-row gap-2">
            <Button
              onClick={() => router.push(`/stock-logs?stock_id=${row.ID}`)}
              variant={'outline'}
              className="rounded-2xl"
            >
              <Eye />
            </Button>
          </div>
        );
      },
    },
  ];

  const fetchData = async () => {
    const params: Record<string, string | number> = { page, pageSize, search };
    if (selectedWarehouseId) {
      params.warehouse_id = selectedWarehouseId;
    }
    const res = await http.get('/api/stocks', { params });
    setData(res.data.data);
    setTotal(res.data.pagination?.total ?? res.data.data.length);
  };

  useEffect(() => {
    const timeout = setTimeout(() => {
      fetchData();
    }, 300);

    return () => clearTimeout(timeout);
  }, [page, search, selectedWarehouseId]);

  return (
    <div className="p-4">
      <div className="flex flex-row items-center justify-between mb-4">
        <div className="flex items-center gap-2">
          <label htmlFor="warehouse-filter" className="text-sm font-medium">
            Filter by Warehouse:
          </label>
          <select
            id="warehouse-filter"
            value={selectedWarehouseId}
            onChange={(e) => {
              setSelectedWarehouseId(e.target.value);
              setPage(1);
            }}
            className="border rounded px-2 py-1 text-sm"
          >
            <option value="">All Warehouses</option>
            {warehouses.map((w) => (
              <option key={w.ID} value={String(w.ID)}>
                {w.name}
              </option>
            ))}
          </select>
        </div>
        <Button
          onClick={() => router.push('/inventory/transaction')}
          className="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition"
        >
          + Transaction
        </Button>
      </div>
      <DataTable
        onSearchChange={setSearch}
        columns={columns}
        data={data}
        page={page}
        pageSize={pageSize}
        total={total}
        onPageChange={setPage}
        search={search}
      />
    </div>
  );
}
