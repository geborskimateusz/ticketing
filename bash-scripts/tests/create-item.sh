#!/bin/bash

echo "Creating new item."
curl -b ./cookie.txt --insecure --header "Content-Type: application/json" \
  --request POST \
  --data '{"title":"Lorem Ipsum","description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Fusce quis nibh nec tellus tempor feugiat. Duis feugiat id lorem non vulputate. Vestibulum scelerisque, nisi non fermentum dictum, ex mi imperdiet est, non finibus sem quam sit amet sapien. Quisque sed urna ornare, posuere turpis id, varius lacus. Maecenas in massa eu quam tristique maximus vel mattis justo. Etiam dapibus, purus in commodo tincidunt, risus mi finibus nibh, in eleifend ipsum sem sit amet quam. Nulla eu condimentum arcu. Cras at enim vel quam feugiat scelerisque id ac augue. Quisque sem dui, vulputate finibus rutrum vitae, malesuada a justo. Vestibulum velit enim, pellentesque vitae sagittis et, ornare sit amet quam. Etiam.", "condition": "new", "quantity": "1", "price": "5"}' \
  https://ticketing.dev/api/items | jq '.'

