import { h } from "vue";
import { NButton } from "naive-ui";
export default function createColumns({click,del}) {
    return [
    {
        title: 'Maintenance Process ID',
        key: 'MaintenanceProcessID'
        },
        {
        title: 'Hardware ID',
        key: 'HardwareID'
        },
        // {
        // title: 'Issue Description',
        // key: 'IssueDescription'
        // },
        // {
        // title: 'Solution Description',
        // key:'SolutionDescription'
        // },
        {
        title: 'Maintenance Date',
        key: 'MaintenanceDate'
        },
        {
        title: 'Cost',
        key: 'Cost'
        },
        {
        title: 'Status',
        key: 'Status'
        },
      {
        title: 'Action',
        key: 'actions',
        render (row) {
          return h(
          'div',
          [
            h(
            NButton,
            {
              strong: true,
              tertiary: true,
              size: 'small',
              onClick: () => click(row)
            },
            { default: () => 'Detail' }
            ),
            h(
              NButton,
              {
                strong: true,
                type: 'error',
                size: 'small',
                onClick: () => del(row)
              },
              { default: () => 'Delete' }
              )
          ]
          
          )
        }
      }
    ]
  }
