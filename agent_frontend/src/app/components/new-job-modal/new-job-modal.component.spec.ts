import { ComponentFixture, TestBed } from '@angular/core/testing';

import { NewJobModalComponent } from './new-job-modal.component';

describe('NewJobModalComponent', () => {
  let component: NewJobModalComponent;
  let fixture: ComponentFixture<NewJobModalComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ NewJobModalComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(NewJobModalComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
