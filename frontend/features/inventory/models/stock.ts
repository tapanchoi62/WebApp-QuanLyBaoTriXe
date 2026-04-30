export interface Stock {
  ID: number;
  item_id: number;
  warehouse_id: number;
  quantity: number;
  Item: { ID: number; name: string; category: string; unit: string };
  Warehouse: { ID: number; name: string; location: string };
  CreatedAt: string;
  UpdatedAt: string;
}
