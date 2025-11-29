"use client";

import { useEffect, useState } from "react";
import { ColumnDef } from "@tanstack/react-table";
import { DataTable } from "@/components/DataTable";
import http from "@/lib/axios";
import { useHeader } from "@/contexts/HeaderContext";
import { Vehicle } from "../models/vehicle";
import { useRouter } from "next/navigation";
import { Button } from "@/components/ui/button";
import { Eye, Trash2 } from "lucide-react";
import { toast } from "sonner";
import { ConfirmButton } from "@/components/ConfirmButton";


export default function VehicleContainer() {
  const [data, setData] = useState<Vehicle[]>([]);
  const [page, setPage] = useState(1);
  const pageSize = 4;
  const [total, setTotal] = useState(0);
  const [search, setSearch] = useState("");
  const router = useRouter();
  const {setTitle} = useHeader();

  useEffect(() => {
    setTitle("Vehicle Management");
  }, [setTitle]);



  const columns: ColumnDef<Vehicle>[] = [
    {
      header: 'ID',
      accessorKey: 'ID',
    },
    {
      header: 'Biển số',
      accessorKey: 'plate_number',
    },
    {
      header: 'Năm SX',
      accessorKey: 'year',
    },
    {
      header: 'Ghi chú',
      accessorKey: 'note',
    },
    {
      header: 'Model',
      accessorKey: 'model_vehicle',
    },
    {
      header: 'Ngày tạo',
      accessorKey: 'CreatedAt',
      cell: info => new Date(info.getValue() as string).toLocaleString(), // format ngày giờ
    },
    {
      header: 'Ngày cập nhật',
      accessorKey: 'UpdatedAt',
      cell: info => {
        const val = info.getValue() as string;
        return val === "0001-01-01T00:00:00Z" ? '-' : new Date(val).toLocaleString();
      },
    },
    {
      header: 'Actions',
      cell: info => {
        const row = info.row.original;
        return (
          <div className="flex flex-row gap-2">
            <Button
              onClick={() => router.push(`/vehicles/${row.ID}`)}
              variant={'outline'}
              className="rounded-2xl"
            >
              <Eye></Eye>
            </Button>
            <ConfirmButton
              variant="destructive"
              dialogTitle="Delete Vehicle"
              dialogDescription="Are you sure you want to delete this vehicle? This action cannot be undone."
              confirmText="Delete"
              cancelText="Cancel"
              onConfirm={() => removeVehicle(row.ID)}
            >
              <Trash2></Trash2>
            </ConfirmButton>
          </div>
          
        );
      }
    }
  ];

  const removeVehicle = async (id: number) => {
    try {
      await http.delete(`/vehicles/${id}`);
      toast("Vehicle deleted successfully");
      fetchData();
    } catch (error) {
      console.error("Delete vehicle failed:", error);
      toast("Failed to delete vehicle");
    }
  };
  const fetchData = async () => {
    const res = await http.get("/vehicles", {
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
    router.push("/vehicles/create");
  }
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
