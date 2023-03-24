import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DashboardNavlinkComponent } from './dashboard-navlink.component';

describe('DashboardNavlinkComponent', () => {
  let component: DashboardNavlinkComponent;
  let fixture: ComponentFixture<DashboardNavlinkComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ DashboardNavlinkComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(DashboardNavlinkComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
