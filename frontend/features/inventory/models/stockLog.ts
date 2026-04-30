export interface StockLog {
  ID: number;
  stock_id: number;
  type: 'IN' | 'OUT' | 'ADJUST';
  quantity: number;
  note: string;
  supplier_id?: number;
  Supplier?: { ID: number; name: string };
  Stock?: {
    Item: { name: string; unit: string };
    Warehouse: { name: string };
  };
  CreatedAt: string;
}
