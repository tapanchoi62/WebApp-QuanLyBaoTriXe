"use client";

import VehicleForm from "@/features/vehicles/componets/VehicleForm";
import { useParams } from "next/navigation";

export default function Page() {
  const params = useParams();
  const idParam = Array.isArray(params.id) ? params.id[0] : params.id;

  return <VehicleForm id={idParam} />;
}