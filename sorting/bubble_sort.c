#include <stdio.h>

int main() {
  int unsorted_array[] = {43, 62, 18, 24, 87, 41, 20, 89, 76, 82};
  int array_size = sizeof(unsorted_array) / sizeof(int);

  int i, j, tmp, swap = 0;
  for (int i = 0; i < array_size; i++) {
    for (int j = 0; j < (array_size - 1); j++) {
      if (unsorted_array[j] > unsorted_array[j+1]) {
        tmp = unsorted_array[j];
        unsorted_array[j] = unsorted_array[j+1];
        unsorted_array[j+1] = tmp;
        swap = 1;
      }

      if (swap == 0) {
        printf("Array already sorted.\n");
        return 0;
      }
    }
  }

  // Print array.
  for (int i = 0; i < array_size; i++) {
    printf("%d ", unsorted_array[i]);
  }
}
