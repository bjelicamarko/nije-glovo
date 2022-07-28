import { Routes } from "@angular/router";
import { RoleGuard } from "../auth/role-guard/role.guard";
import { OrdersPageComponent } from "../shared/pages/orders-page/orders-page.component";

export const EmployeeRoutes: Routes = [
    {
        path: "orders",
        pathMatch: "full",
        component: OrdersPageComponent,
        canActivate: [RoleGuard],
        data: { expectedRoles: "EMPLOYEE" },
    }
];