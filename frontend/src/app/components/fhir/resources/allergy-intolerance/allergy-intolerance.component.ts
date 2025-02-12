import {ChangeDetectorRef, Component, Input, OnInit} from '@angular/core';
import {FhirResourceComponentInterface} from '../../fhir-resource/fhir-resource-component-interface';
import {TableRowItem, TableRowItemDataType} from '../../common/table/table-row-item';
import {Router, RouterModule} from '@angular/router';
import {AllergyIntoleranceModel} from '../../../../../lib/models/resources/allergy-intolerance-model';
import {NgbCollapseModule} from "@ng-bootstrap/ng-bootstrap";
import {CommonModule} from "@angular/common";
import {BadgeComponent} from "../../common/badge/badge.component";
import {TableComponent} from "../../common/table/table.component";

@Component({
  standalone: true,
  imports: [NgbCollapseModule, CommonModule, BadgeComponent, TableComponent, RouterModule],
  selector: 'fhir-allergy-intolerance',
  templateUrl: './allergy-intolerance.component.html',
  styleUrls: ['./allergy-intolerance.component.scss']
})
export class AllergyIntoleranceComponent implements OnInit, FhirResourceComponentInterface {
  @Input() displayModel: AllergyIntoleranceModel
  @Input() showDetails: boolean = true

  isCollapsed: boolean = false

  tableData: TableRowItem[] = []

  constructor(public changeRef: ChangeDetectorRef, public router: Router) {}

  ngOnInit(): void {
    this.tableData = [
      {
        label: 'Substance',
        data: this.displayModel?.code,
        data_type: TableRowItemDataType.CodableConcept,
        enabled: !!this.displayModel?.code,
      },
      {
        label: 'Type',
        data: this.displayModel?.type,
        enabled: !!this.displayModel?.type,
      },
      {
        label: 'Category',
        data: this.displayModel?.category?.join(" "),
        enabled: !!this.displayModel?.category,
      },
    {
      label: 'Patient',
      data: this.displayModel?.patient,
      data_type: TableRowItemDataType.Reference,
      enabled: !!this.displayModel?.patient,
    },
    {
      label: 'Asserted by',
      data: this.displayModel?.asserter,
      data_type: TableRowItemDataType.Reference,
      enabled: !!this.displayModel?.asserter,
    },
    // {
    //   label: 'Manifestation',
    //     testId: 'manifestation',
    //   data: reaction.map((reaction, i) => {
    //   const manifestations = _get(reaction, 'manifestation', []);
    //   const severity = _get(reaction, 'severity');
    //   return manifestations.map((manifestation, j) => {
    //     return (
    //       <div key={`item-${i}${j}`} className="d-flex">
    //     <CodeableConcept fhirData={manifestation} />
    //     {severity && (
    //       <span className="ms-4">
    //         <BadgeSecondary>{severity}</BadgeSecondary>
    //         </span>
    //     )}
    //     </div>
    //   );
    //   });
    // }),
    //   status: hasReaction,
    // },
    // {
    //   label: 'Notes',
    //   data: hasNote && <Annotation fhirData={note} />,
    //   status: this.displayModel.has_nod,
    // }
    ];
  }
  markForCheck(){
    this.changeRef.markForCheck()
  }
}
