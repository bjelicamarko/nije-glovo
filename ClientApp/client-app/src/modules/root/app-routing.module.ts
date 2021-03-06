import { NgModule } from "@angular/core";
import { Routes, RouterModule } from "@angular/router";
import { RoleGuard } from "../auth/role-guard/role.guard";
import { UserInfoComponent } from "../shared/pages/user-info/user-info.component";
import { MainPageComponent } from "./pages/main-page/main-page.component";
import { NotFoundPageComponent } from "./pages/not-found-page/not-found-page.component";
import { RegistrationPageComponent } from "./pages/registration-page/registration-page.component";
import { RootLayoutPageComponent } from "./pages/root-layout-page/root-layout-page.component";

const routes: Routes = [
  {
    path: "app",
    component: RootLayoutPageComponent,
    children: [
      {
        path: "auth",
        loadChildren: () =>
          import("./../auth/auth.module").then((m) => m.AuthModule),
      },
      {
        path: "registration",
        component: RegistrationPageComponent
      },
      {
        path: "main",
        component: MainPageComponent,
        children: [
          {
            path: "admin",
            loadChildren: () =>
              import("./../admin/admin.module").then((m) => m.AdminModule),
          },
          {
            path: "appuser",
            loadChildren: () =>
              import("./../appuser/appuser.module").then((m) => m.AppUserModule)
          },
          {
            path: "deliverer",
            loadChildren: () =>
              import("./../deliverer/deliverer.module").then((m) => m.DelivererModule)
          },
          {
            path: "employee",
            loadChildren: () =>
              import("./../employee/employee.module").then((m) => m.EmployeeModule)
          },
          {
            path: "profile",
            component: UserInfoComponent,
            canActivate: [RoleGuard],
            data: { expectedRoles: "ADMIN|APPUSER|DELIVERER|EMPLOYEE" },
          },
        ]
      }
    ]
  },
  {
    path: "",
    redirectTo: "app/auth/login",
    pathMatch: "full",
  },
  {
    path: "**",
    component: NotFoundPageComponent
  },
];
  
@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }