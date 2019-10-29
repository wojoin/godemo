#ifndef __EXAMPLE_H__
#define __EXAMPLE_H__

typedef struct {
  int x,y;
} Point;

enum color { RED, BLUE, GREEN };

class Foo {
 public:
  Foo() { }
  enum speed { IMPULSE=10, WARP=20, LUDICROUS=30 };
  void enum_test(speed s);
};

void enum_test(color c, Foo::speed s);


class Shape {
public:
  Shape() {
    nshapes++;
  }
  virtual ~Shape() {
    nshapes--;
  }
  double  x, y;	// center point
  void    move(double dx, double dy);
  virtual double area() = 0;
  virtual double perimeter() = 0;
  static  int nshapes;
};

class Circle : public Shape {
private:
  double radius;
public:
  Circle(double r) : radius(r) { }
  virtual double area();
  virtual double perimeter();
};

class Square : public Shape {
private:
  double width;
public:
  Square(double w) : width(w) { }
  virtual double area();
  virtual double perimeter();
};

#endif