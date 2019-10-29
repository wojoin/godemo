/* File : example.c */

/* A global variable */

#include <stdio.h>
#include <stdlib.h>
#include "example.h"


#define M_PI 3.14159265358979323846

double 	Dvar = 3.0;
int 	status = 1;
char 	path[256] = "/home/join";
int     ivar = 42;
int*	iptrvar = &ivar;

/* Global variables involving a structure */
Point*	ptptr = 0;
Point 	pt = { 10, 20 };


void print_vars() {
	printf("print info from C\n");
	printf("status		= %d\n", status);
	printf("path		= %s\n", path);
	
	printf("iptrvar   	= %p\n", iptrvar);
  	
	//printf("ptptr     = %p (%d, %d)\n", ptptr, ptptr ? ptptr->x : 0, ptptr ? ptptr->y : 0);
	printf("point       = (%d, %d)\n", pt.x, pt.y);
  
}


char* point_Print(Point *p) {
  static char buffer[256];
  if (p) {
    sprintf(buffer,"(%d,%d)", p->x,p->y);
  } else {
    sprintf(buffer,"null");
  }
  return buffer;
}


/* A function to create an integer (to test iptrvar) */
int* new_int(int value) {
  int *ip = (int *) malloc(sizeof(int));
  *ip = value;
  return ip;
}

/* A function to create a point */
Point* new_Point(int x, int y) {
  Point *p = (Point *) malloc(sizeof(Point));
  p->x = x;
  p->y = y;
  return p;
}

/* Compute the greatest common divisor of positive integers */
int gcd(int x, int y) {
      int g;
      g = y;
      while (x > 0) {
         g = x;
         x = y % x;
         y = g;
      }
      return g;
}

/* enum */
void Foo::enum_test(speed s) {
  if (s == IMPULSE) {
    printf("IMPULSE speed\n");
  } else if (s == WARP) {
    printf("WARP speed\n");
  } else if (s == LUDICROUS) {
    printf("LUDICROUS speed\n");
  } else {
    printf("Unknown speed\n");
  }
}


void enum_test(color c, Foo::speed s) {
  if (c == RED) {
    printf("color = RED, ");
  } else if (c == BLUE) {
    printf("color = BLUE, ");
  } else if (c == GREEN) {
    printf("color = GREEN, ");
  } else {
    printf("color = Unknown color!, ");
  }

  if (s == Foo::IMPULSE) {
    printf("speed = IMPULSE speed\n");
  } else if (s == Foo::WARP) {
    printf("speed = WARP speed\n");
  } else if (s == Foo::LUDICROUS) {
    printf("speed = LUDICROUS speed\n");
  } else {
    printf("speed = Unknown speed!\n");
  }
}


/* class */
int Shape::nshapes = 0;

/* Move the shape to a new location */
void Shape::move(double dx, double dy) {
  x += dx;
  y += dy;
}

double Circle::area() {
  return M_PI * radius * radius;
}

double Circle::perimeter() {
  return 2 * M_PI * radius;
}

double Square::area() {
  return width * width;
}

double Square::perimeter() {
  return 4 * width;
}
