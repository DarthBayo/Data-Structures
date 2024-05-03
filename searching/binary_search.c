#include <stdio.h>
#include <math.h>

int main() {
  int sorted_array[] = {8, 13, 20, 21, 30, 33, 49, 50, 52, 64}, value = 50;
  int array_size = sizeof(sorted_array) / sizeof(int);
  int low = 0, mid = 0, high = array_size;

  // Print array.
  for (int i = 0; i < array_size; i++) {
    printf("%d ", sorted_array[i]);
  }

  for (int i = 0; i < array_size; i++) {
    mid = low + floor((high - low) / 2);
    if (sorted_array[mid] == value) {
      printf("\nValue: %d\nArray size: %d\nPosition: %d\nLoops: %d", value, array_size, mid, (i + 1));
      return 0;
    }

    if (sorted_array[mid] > value) {
      high = mid - 1;
    }
    else {
      low = mid + 1;
    }
  }
}
