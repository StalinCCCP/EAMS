import { h } from "vue";
import { NButton } from "naive-ui";
export default function createColumns({click,del}) {
    return [
      {
        title: 'HardwareID',
        key: 'HardwareID'
      },
      {
        title: 'Hardware Name',
        key: 'HardwareName'
      },
      {
        title: 'Category',
        key: 'Category'
      },
      {
        title: 'Status',
        key: 'Status'
      },
      {
        title: 'Location',
        key: 'Location'
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
