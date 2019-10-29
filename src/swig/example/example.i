/* File : example.i */
%module example

%{
#include "example.h"
%}

#define PI 3.14159265358979323846
#define ICONST 42

#define SCONST "Hello World"

#define AREA PI * ICONST * ICONST
#define LENGTH (2 * (ICONST) * PI)

%constant int iconst = 37;
%constant double fconst = 3.14;



%inline %{
extern int    	gcd(int x, int y);
extern double 	Dvar;

extern int  	status;
extern char 	path[256];
extern void  	print_vars();

extern int*		iptrvar;
extern int* 	new_int(int value);

extern Point*	ptptr;
extern Point* 	new_Point(int x, int y);
extern char* 	point_Print(Point *p);

%}

/* Let's just grab the original header file here */
%nodefaultctor Point;
%nodefaultdtor Point;

%include "example.h"