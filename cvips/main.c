#include <stdio.h>
#include <vips/vips.h>

void resize_step() {
  VipsImage *in;
  VipsImage *out;

  if( !(in = vips_image_new_from_file( "images/image50.jpg", NULL )) )
    vips_error_exit( "Error while loading file" );

  if( vips_resize( in, &out, 0.57550644567, NULL ) )
    vips_error_exit( "Error resizing" );

  g_object_unref( in ); 

  if( vips_image_write_to_file( out, "out/resize.jpg", NULL ) )
    vips_error_exit( "Error while writing file" );

  g_object_unref( out ); 
}

void resize( int argc, char **argv ) {
   int iterations;

   if ( argc < 3 ) {
      vips_error_exit( "Please provide the number of iterations.\n" );
   }

   iterations = atoi(argv[2]);

   printf("resize with %d iterations\n", iterations);

   for (int i = 0; i < iterations; i++) {
      printf("iteration %d\n", i);
      resize_step();
   }
}

int main( int argc, char **argv ) {
   printf("\n");
   
   if( VIPS_INIT( argv[0] ) )
      vips_error_exit( NULL );

   if ( argc < 2 ) {
      vips_error_exit( "Please provide the name of the test as argument.\n" );
   }

   printf("Running test %s\n", argv[1]);

   if ( strcmp(argv[1], "resize") == 0 ) {
      resize( argc, argv );
   } else {
      vips_error_exit( "Unknown test.\n" );
   }

   vips_shutdown();

  return( 0 );
}

