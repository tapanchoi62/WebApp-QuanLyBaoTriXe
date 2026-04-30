'use client';

import { useEffect, useState } from 'react';
import { ColumnDef } from '@tanstack/react-table';
import { DataTable } from '@/components/DataTable';
import http from '@/lib/axios';
import { useHeader } from '@/contexts/HeaderContext';
import { StockLog } from '../models/stockLog';
import { Warehouse } from '../models/warehouse';
import { useSearchParams } from 'next/navigation';

export default function StockLogContainer() {
  const [data, setData] = useState<StockLog[]>([]);
  const [page, setPage] = useState(1);
  const pageSize = 20;
  const [total, setTotal] = useState(0);
  const [search, setSearch] = useState('');
  const [warehouses, setWarehouses] = useState<Warehouse[]>([]);
  const [selectedWarehouseId, setSelectedWarehouseId] = useState<string>('');
  const { setTitle } = useHeader();
  const searchParams = useSearchParams();

  useEffect(() => {
    setTitle('Stock Logs');
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

  const columns: ColumnDef<StockLog>[] = [
    {
      header: 'ID',
      accessorKey: 'ID',
    },
    {
      header: 'Type',
      accessorKey: 'type',
      cell: (info) => {
        const row = info.row.original;
        return (
          <span
            className={`px-2 py-1 rounded text-xs font-medium ${
              row.type === 'IN'
                ? 'bg-green-100 text-green-800'
                : row.type === 'OUT'
                  ? 'bg-red-100 text-red-800'
                  : 'bg-yellow-100 text-yellow-800'
            }`}
          >
            {row.type}
          </span>
        );
      },
    },
    {
      header: 'Quantity',
      accessorKey: 'quantity',
    },
    {
      header: 'Note',
      accessorKey: 'note',
    },
    {
      header: 'Supplier',
      accessorKey: 'Supplier',
      cell: (info) => {
        const row = info.row.original;
        return row.Supplier?.name ?? '-';
      },
    },
    {
      header: 'Created At',
      accessorKey: 'CreatedAt',
      cell: (info) => new Date(info.getValue() as string).toLocaleString(),
    },
  ];

  const fetchData = async () => {
    const params: Record<string, string | number> = { page, pageSize, search };
    if (selectedWarehouseId) {
      params.warehouse_id = selectedWarehouseId;
    }
    const stockId = searchParams.get('stock_id');
    if (stockId) {
      params.stock_id = stockId;
    }
    const res = await http.get('/api/stock-logs', { params });
    setData(res.data.data);
    setTotal(res.data.pagination?.total ?? res.data.data.length);
  };

  useEffect(() => {
    const timeout = setTimeout(() => {
      fetchData();
    }, 300);

    return () => clearTimeout(timeout);
  }, [page, search, selectedWarehouseId, searchParams]);

  return (
    <div className="p-4">
      <div className="flex items-center gap-2 mb-4">
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
