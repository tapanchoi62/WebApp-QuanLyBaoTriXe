"use client";

import {
  flexRender,
  getCoreRowModel,
  getSortedRowModel,
  useReactTable,
  ColumnDef,
} from "@tanstack/react-table";
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import DataTablePagination from "./DataTablePaging";
import DataTableToolbar from "./DataTableToolbar";
import { Button } from "./ui/button";

interface DataTableProps<TData, TValue> {
  columns: ColumnDef<TData, TValue>[];
  data: TData[];
  page: number;
  pageSize: number;
  total: number;
  onPageChange: (page: number) => void;
  onSearchChange: (v: string) => void;
  onCreate?: () => void;
  search: string;
}

export function DataTable<TData, TValue>({
  columns,
  data,
  page,
  pageSize,
  total,
  onPageChange,
  onCreate,
  onSearchChange,
  search
}: DataTableProps<TData, TValue>) {
  const table = useReactTable({
    data,
    columns,
    state: {},
    getCoreRowModel: getCoreRowModel(),
    getSortedRowModel: getSortedRowModel(),
  });

  return (
    <div className="border rounded-md px-4">
      <div className="flex flex-row items-start justify-between w-full">
        <DataTableToolbar search={search} onSearchChange={onSearchChange}/>
        {onCreate && (
          <div className="flex justify-end py-2">
            <Button
              onClick={onCreate}
              className="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition"
            >
              Create New
            </Button>
          </div>
        )}
      </div>
      
      <Table>
        <TableHeader>
          {table?.getHeaderGroups().map((hg) => (
            <TableRow key={hg.id}>
              {hg.headers.map((h) => (
                <TableHead
                  key={h.id}
                  onClick={h.column.getToggleSortingHandler()}
                  className="cursor-pointer select-none"
                >
                  {flexRender(h.column.columnDef.header, h.getContext())}
                  {h.column.getIsSorted() === "asc" ? " 🔼" : ""}
                  {h.column.getIsSorted() === "desc" ? " 🔽" : ""}
                </TableHead>
              ))}
            </TableRow>
          ))}
        </TableHeader>

        <TableBody>
          {table?.getRowModel().rows.map((row) => (
            <TableRow key={row.id}>
              {row.getVisibleCells().map((cell) => (
                <TableCell key={cell.id}>
                  {flexRender(cell.column.columnDef.cell, cell.getContext())}
                </TableCell>
              ))}
            </TableRow>
          ))}
        </TableBody>
      </Table>

      <DataTablePagination
        page={page}
        pageSize={pageSize}
        total={total}
        onPageChange={onPageChange}
      />
    </div>
  );
}

export default DataTable;
