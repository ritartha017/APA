// win
// #include <stdio.h>
// #include <windows.h>

// void print_a(){
// 	while(true){
// 		Sleep(500);
// 		printf("A");
// 	}
// }

// void print_b(){
// 	while(true){
// 		Sleep(500);
// 		printf("B");
// 	}
// }

// int main(){
// 	unsigned thread_id;
// 	HANDLE handler;
// 	handler = CreateThread(NULL, 0, print_a, NULL, 0, &thread_id);
// 	printf("thread_id is %u", thread_id);
// 	print_b();
// }

// linux
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>  //Header file for sleep(). man 3 sleep for details.
#include <pthread.h>

void *print_a() {
	int i = 0;
    while(i++ < 4) {
		sleep(1);
		printf("A");
		fflush(stdout); // Force the output to be printed
	}
	printf("A finish\n");
}

void *print_b() {
	int i = 0;
    while(i++ < 4) {
		sleep(1);
		printf("B");
		fflush(stdout);
	}
	printf("B finish\n");
}

int main() {
    pthread_t thread_id1, thread_id2;

    pthread_create(&thread_id1, NULL, print_a, NULL);
    printf("1 thread_id is %lu\n", thread_id1);
    pthread_join(thread_id1, NULL);

    pthread_create(&thread_id2, NULL, print_b, NULL);
    printf("2 thread_id is %lu\n", thread_id2);
    pthread_join(thread_id2, NULL);
}