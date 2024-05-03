#include <stdio.h>

int main() {
  int unsorted_array[] = {43, 62, 18, 24, 87, 41, 20, 89, 76}, value = 20;
  int array_size = sizeof(unsorted_array) / sizeof(int);

  // Print array.
  for (int i = 0; i < array_size; i++) {
    printf("%d ", unsorted_array[i]);
  }

  for (int i = 0; i < array_size; i++) {
    if (unsorted_array[i] == value) {
      printf("\nValue: %d\nPosition: %d\nLoops: %d", value, i, (i + 1));
    }
  }
}
