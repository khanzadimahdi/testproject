import {TABLE_HEADERS} from "./comments-table";
import {TableSkeleton} from "../table-skeleton";

export function CommentsTableSkeleton() {
  return <TableSkeleton headers={TABLE_HEADERS} />;
}