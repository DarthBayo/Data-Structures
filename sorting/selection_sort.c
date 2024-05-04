#include <stdio.h>

int main() {
  int unsorted_array[] = {43, 62, 18, 24, 87, 41, 20, 89, 76, 82};
  int array_size = sizeof(unsorted_array) / sizeof(int);

  int i, j, tmp, swap = 0, idx = 0;
  for (int i = 0; i < array_size; i++) {
    idx = i;
    for (int j = i; j < array_size; j++) {
      if (unsorted_array[idx] > unsorted_array[j]) {
        idx = j;
        swap = 1;
      }
    }

    if (swap == 0) {
      printf("Array already sorted.\n");
      return 0;
    }

    if (i != idx) {
      tmp = unsorted_array[idx];
      unsorted_array[idx] = unsorted_array[i];
      unsorted_array[i] = tmp;
    }
  }

  // Print array.
  for (int i = 0; i < array_size; i++) {
    printf("%d ", unsorted_array[i]);
  }
}
